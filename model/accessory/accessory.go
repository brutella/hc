package accessory

import(
    "github.com/brutella/hap/model/service"
)

type Accessory struct {
    Id int              `json:"aid"`
    Services []*service.Service `json:"services"`
    
    idCount int
}

func NewAccessory() *Accessory {
    return &Accessory{
        idCount: 1,
    }
}

// Adds a service to the accessory and updates the ids of the service and the corresponding characteristics
func (a *Accessory) AddService(s *service.Service) {
    s.Id = a.idCount
    a.idCount += 1
    
    for _, c := range s.Characteristics {
        c.Id = a.idCount
        a.idCount += 1
    }
    
    a.Services = append(a.Services, s)
}

func (a *Accessory) Equal(other interface{}) bool {
    if accessory, ok := other.(*Accessory); ok == true {
        if len(a.Services) != len(accessory.Services) {
            return false
        }
        
        for i, s := range a.Services {
            if s.Equal(accessory.Services[i]) == false {
                return false
            }
        }
        
        return a.Id == accessory.Id
    }
    
    return false
}
