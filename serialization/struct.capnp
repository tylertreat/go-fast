@0x8b0d616b5401e931;

using Go = import "go.capnp";
$Go.package("serialization");

struct StructCapnp {
    field1 @0 :Text;
    field2 @1 :Int64;
    field3 @2 :List(Text);
    field4 @3 :UInt64;
    field5 @4 :Text;
    field6 @5 :Text;
    field7 @6 :Data;
}
