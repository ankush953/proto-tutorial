package command

import (
	"fmt"

	"github.com/ankush953/proto-tutorial.git/proto/generated_code"
)

func OneofDemo() {
	fmt.Println("Welcome to demo of oneof feature in protobuf...")
	pan := "THISISPAN"
	zip := "THISISZIP"
	aadhar := &generated_code.OneofDemo_Aadhar{
		Aadhar: "THISISAADHAR",
	}
	registration := &generated_code.OneofDemo_Registration{
		Registration: "THISISREGISTRATION",
	}
	demo := &generated_code.OneofDemo{
		Pan:                  &pan,
		Zip:                  &zip,
		AadharOrRegistration: &generated_code.OneofDemo_Aadhar{},
	}
	aadharOrRegistrationDescriptor := demo.ProtoReflect().Descriptor().Oneofs().ByName("aadhar_or_registration")
	field := demo.ProtoReflect().WhichOneof(aadharOrRegistrationDescriptor)
	fmt.Printf("Currently set Oneof field: %v, aadhar: %v, registration: %v\n", field.FullName(), demo.GetAadhar(), demo.GetRegistration())
	demo.AadharOrRegistration = aadhar
	field = demo.ProtoReflect().WhichOneof(aadharOrRegistrationDescriptor)
	fmt.Printf("Currently set Oneof field: %v, aadhar: %v, registration: %v\n", field.FullName(), demo.GetAadhar(), demo.GetRegistration())
	demo.AadharOrRegistration = registration
	field = demo.ProtoReflect().WhichOneof(aadharOrRegistrationDescriptor)
	fmt.Printf("Currently set Oneof field: %v, aadhar: %v, registration: %v\n", field.FullName(), demo.GetAadhar(), demo.GetRegistration())
}
