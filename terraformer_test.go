package ogame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTerraformerIsAvailable(t *testing.T) {
	resourcesBuildings := ResourcesBuildings{}
	facilities := Facilities{NaniteFactory: 1}
	researches := Researches{EnergyTechnology: 12}
	b := newTerraformer()
	avail := b.IsAvailable(PlanetDest, resourcesBuildings, facilities, researches, 0)
	assert.True(t, avail)
}

func TestTerraformerIsAvailable_NoTech(t *testing.T) {
	resourcesBuildings := ResourcesBuildings{}
	facilities := Facilities{NaniteFactory: 1}
	researches := Researches{EnergyTechnology: 11}
	b := newTerraformer()
	avail := b.IsAvailable(PlanetDest, resourcesBuildings, facilities, researches, 0)
	assert.False(t, avail)
}
