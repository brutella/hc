package model

type HeatingCoolingMode byte

const (
    // TODO verify the values
    ModeOff = HeatingCoolingMode(0x00)
    ModeHeating = HeatingCoolingMode(0x01)
    ModeCooling = HeatingCoolingMode(0x02)
)