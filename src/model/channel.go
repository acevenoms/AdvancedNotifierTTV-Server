package model

type Channel struct {
	Name    string
	Viewers []Viewer
}

func (c *Channel) AddViewer(newV Viewer) *Viewer {
	for i, oldV := range c.Viewers {
		if oldV.Name == newV.Name {
			return &c.Viewers[i]
		}
	}

	c.Viewers = append(c.Viewers, newV)
	return &newV
}

func (c Channel) GetViewer(userId string) *Viewer {
	for i, v := range c.Viewers {
		if v.Name == userId {
			return &c.Viewers[i]
		}
	}

	return nil
}
