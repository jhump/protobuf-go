package protoreflect_test

import (
	"testing"

	"google.golang.org/protobuf/internal/testprotos/test3"
	"google.golang.org/protobuf/proto"
)

func TestSyntheticOneof(t *testing.T) {
	msg := test3.TestAllTypes{}
	md := msg.ProtoReflect().Descriptor()
	ood := md.Oneofs().ByName("_optional_int32")
	if ood == nil {
		t.Fatal("failed to find oneof _optional_int32")
	}
	if !ood.IsSynthetic() {
		t.Fatal("oneof _optional_int32 should be synthetic")
	}
	if msg.ProtoReflect().WhichOneof(ood) != nil {
		t.Error("oneof _optional_int32 should not have a field set yet")
	}
	msg.OptionalInt32 = proto.Int32(123)
	if msg.ProtoReflect().WhichOneof(ood) == nil {
		t.Error("oneof _optional_int32 should have a field set")
	}
}
