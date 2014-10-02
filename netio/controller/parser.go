package controller

import(
    "strings"
    "strconv"
    "github.com/brutella/hap/common"
)

// string must be in format <accessory id>.<characteristic id>
func ParseAccessoryAndCharacterId(str string) (int, int, error) {
    ids := strings.Split(str, ".")
    if len(ids) != 2 {
        return 0, 0, common.NewErrorf("Could not parse uid %s", str)
    }
    
    aid, err := strconv.Atoi(ids[0])
    cid, err := strconv.Atoi(ids[1])
    
    return aid, cid, err
}