package model

import "log"


type modelMetrics struct {
    delegate Model
}

func newMetrics(m Model) (Model, error) {
    return &modelMetrics{
        delegate: m
    }, nil
}

func onStart(nm string) {
    log.Printf("%s started", nm)
}

func onEnd(name string) {
    log.Printf("%s ended", nm)
}
