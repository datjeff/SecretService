package models
import(
	"errors"
)

type SecretThingCollection struct {
	secretThings map[int]*SecretThing
	thingCount int
}

func NewSecretThingCollection() *SecretThingCollection {
	return &SecretThingCollection{make(map[int]*SecretThing),0}
}

func (st *SecretThingCollection) AddSecretThing(content string) int {
	var id = st.thingCount
	st.secretThings[id] = &SecretThing{id, content}
	st.thingCount = st.thingCount + 1
	return id
}

func (st *SecretThingCollection) UpdateSecretThing(id int, content string) error {
	if st.secretThings[id] == nil {
		return errors.New("That id does not exist.")
	}
	st.secretThings[id].Content = content
	return nil
}

func (st *SecretThingCollection) RemoveSecretThing(id int) error {
	if st.secretThings[id] == nil {
		return errors.New("That id does not exist.")
	}else{
		delete(st.secretThings,id)
		return nil
	}
}

func (st *SecretThingCollection) GetSecretThing(id int)(*SecretThing, error) {
	if st.secretThings[id] == nil {
		return nil, errors.New("That id does not exist.")
	}else{
		return st.secretThings[id],nil
	}
}

func (st *SecretThingCollection) GetAllSecretThings()[]*SecretThing {
	return nil
}




