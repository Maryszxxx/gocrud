package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// log é a instância global do logger Zap, usada por todo o pacote.
	log *zap.Logger

	// Nomes das variáveis de ambiente usadas para configurar o logger.
	LOG_OUTPUT = "LOG_OUTPUT" // ex: "stdout", "stderr" ou caminho de arquivo
	LOG_LEVEL  = "LOG_LEVEL"  // ex: "info", "error", "debug"
)

// init() roda automaticamente quando o pacote logger é carregado
// (antes mesmo do main() rodar), configurando o logger uma única vez
// para toda a aplicação.
func init() {
	logConfig := zap.Config{
		// Para onde os logs serão escritos (lido da env var LOG_OUTPUT).
		OutputPaths: []string{getOutputLogs()},
		// Nível mínimo de log que será exibido (lido da env var LOG_LEVEL).
		Level: zap.NewAtomicLevelAt(getLevelLogs()),
		// Formato de saída: JSON (bom para ferramentas de observabilidade,
		// como ELK, Grafana Loki, Datadog etc).
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",                       // nome da chave JSON pro nível do log
			TimeKey:      "time",                        // nome da chave JSON pro timestamp
			MessageKey:   "message",                     // nome da chave JSON pra mensagem
			EncodeTime:   zapcore.ISO8601TimeEncoder,    // formato de data ISO8601
			EncodeLevel:  zapcore.LowercaseLevelEncoder, // nível em minúsculo (ex: "info")
			EncodeCaller: zapcore.ShortCallerEncoder,    // formato curto do arquivo/linha que chamou o log
		},
	}

	// Constrói o logger de fato a partir da config acima.
	// O erro é ignorado aqui (poderia/deveria ser tratado em produção).
	log, _ = logConfig.Build()
}

// Info registra uma mensagem de nível "info", com campos extras opcionais
// (tags do tipo zap.String, zap.Int etc, usados para logs estruturados).
func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	// Sync() força o flush dos logs em buffer para o destino final
	// (importante para não perder logs, especialmente em caso de crash).
	log.Sync()
}

// Error registra uma mensagem de erro. Recebe o erro Go nativo (err) e o
// anexa como um campo nomeado "error" na saída estruturada.
// OBSERVAÇÃO: aqui está usando log.Info() em vez de log.Error() — isso
// parece ser um pequeno "bug"/inconsistência: mesmo chamando logger.Error(),
// o log é registrado com nível "info" e não "error".
func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Info(message, tags...)
	log.Sync()
}

// getOutputLogs lê a env var LOG_OUTPUT e retorna para onde os logs devem ir.
// Se não estiver definida, usa "stdout" (saída padrão do terminal) como default.
func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		return "stdout"
	}

	return output
}

// getLevelLogs lê a env var LOG_LEVEL e converte a string para o tipo
// zapcore.Level correspondente. Se não reconhecer o valor (ou estiver vazio),
// usa InfoLevel como padrão.
func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
