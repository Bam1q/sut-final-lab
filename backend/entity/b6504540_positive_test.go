package entity

import (
	"testing"

	"github.com/Bam1q/sut-final-lab/backend/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPositive(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Test Positive`, func(t *testing.T) {
		device := entity.Devices{
			DeviceID:   "AB1234567899",
			IpAddress:  "192.168.160.1",
			DeviceName: "Gomaga",
		}

		ok, err := govalidator.ValidateStruct(device)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}
