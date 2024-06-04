package env

import "log/slog"


type Env struct {
    Logger *slog.Logger
}

func NewEnv() Env {
    return Env{
        Logger: slog.Default(),
    }
}
