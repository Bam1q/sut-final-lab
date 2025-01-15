package entity

import (
	"testing"

	"github.com/Bam1q/sut-final-lab/backend/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNagative(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Test Positive`, func(t *testing.T) {
		device := entity.Devices{
			DeviceID:   "AB1234567899",
			IpAddress:  "192.168", // ผิด
			DeviceName: "Gomaga",
		}

		ok, err := govalidator.ValidateStruct(device)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expext(err.Error()).To(Equal("IpAddress must be a valid IP address"))

	})
}
