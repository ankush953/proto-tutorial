# proto-tutorial
## Table of contents
[Best Practices](#best-pratices)  
[Scalar Type](#scalar-type)  
[Default Field Values](#default-fields)  
[Maps in protocol buffer](#maps)  
[Updating a mesage type](#updating-a-message-type)  
[Extending a Protocol Buffer](#extending-a-protocol-buffer)  
[Exceptions for protocol buffer extension](#exceptions)  
[Extension](#extensions)
[Assigning field numbers](#assining-field-numbers)  
[Reserve field numbers](#reserve-field-numbers)   
[How Message wiring happens in protobuf](#how-protocol-buffers-encode-data)  
[How to move .proto files without much changes](#how-to-move-proto-files-effectively)  
[Nested messages](#nested-messages)  
[OneOf](#oneof-usage)  
[Defining Services](#defining-services)  
[Reading list](#reading-list)

## Best Pratices
Link: https://protobuf.dev/best-practices/dos-donts/  


## Scalar Type
Link: https://protobuf.dev/programming-guides/proto3/#scalar  
You can consider this as native data types supported by Proto. 

## Default Fields
Link: https://protobuf.dev/programming-guides/proto3/#default  
- For strings, the default value is the empty string.
- For bytes, the default value is empty bytes.
- For bools, the default value is false.
- For numeric types, the default value is zero.
- For message fields, the field is not set. Its exact value is language-dependent. See the generated code guide for details.
- For enums, the default value is the first defined enum value, which must be 0.
- In proto3, the first value defined in an enum definition **must** have the value zero and **should** have the name **_ENUM_TYPE_NAME_UNSPECIFIED_** or **_ENUM_TYPE_NAME_UNKNOWN_**.

## Maps
Link: https://protobuf.dev/programming-guides/proto3/#maps  
You can have it as 
```protobuf
map<string, int32> orderToAmount = 6;
```

## Updating a Message Type
Link: https://protobuf.dev/programming-guides/proto3/#updating  
You will need to update the proto message defintions from time to time based on requirements. In this case, Please read below in order
1. [Best Practices](#best-pratices) 
2. [Updating Mesage Type](#updating-a-message-type)
3. [Extending protocol buffer](#extending-a-protocol-buffer)
4. [Exceptions for extension rules](#exceptions)

## Extending a protocol Buffer
Link: https://protobuf.dev/getting-started/gotutorial/#extending-a-protobuf  
Sooner or later after you release the code that uses your protocol buffer, you will undoubtedly want to “improve” the protocol buffer’s definition. If you want your new buffers to be backwards-compatible, and your old buffers to be forward-compatible – and you almost certainly do want this – then there are some rules you need to follow. In the new version of the protocol buffer:

you must not change the tag numbers of any existing fields.
you may delete fields.
you may add new fields but you must use fresh tag numbers (i.e. tag numbers that were never used in this protocol buffer, not even by deleted fields).
(There are some exceptions to these rules, but they are rarely used.)

If you follow these rules, old code will happily read new messages and simply ignore any new fields. To the old code, singular fields that were deleted will simply have their default value, and deleted repeated fields will be empty. New code will also transparently read old messages.

However, keep in mind that new fields will not be present in old messages, so you will need to do something reasonable with the default value. A type-specific default value is used: for strings, the default value is the empty string. For booleans, the default value is false. For numeric types, the default value is zero.

### Exceptions
Link: https://protobuf.dev/programming-guides/proto3/#updating  

## Extensions
Link: https://protobuf.dev/programming-guides/proto2/#extensions  
[NOT SUPPORTED IN PROTO3]
An extension is a field defined outside of its container message; usually in a .proto file separate from the container message’s .proto file.

## Assining Field Numbers
Link: https://protobuf.dev/programming-guides/proto3/#assigning  
- Field number range: `[1 to 536,870,911]` except `[19,000 to 19,999]`
- You should use the field numbers 1 through 15 for the most-frequently-set fields. Lower field number values take less space in the wire format. 
- Field numbers should never be reused
- “Changing” a field number is equivalent to deleting that field and creating a new field with the same type but a new number.

### Reserve Field Numbers 
Link: https://protobuf.dev/programming-guides/proto3/#fieldreserved  
If you have deleted any field, you should reserve its tag number.
```protobuf
reserved 9 to 11;
```
If you use any reserved field number, you would get error
```
proto/addressbook.proto:22:12: Field "phones" uses reserved number 4.
proto/addressbook.proto:22:12: Suggested field numbers for tutorial.Person: 7
```

## How Protocol Buffers encode data 
Link: https://protobuf.dev/programming-guides/encoding/  

### Field presence
Link: https://protobuf.dev/programming-guides/field_presence/

## Specifying Cardinality Labels
Link: https://protobuf.dev/programming-guides/proto3/#field-labels  
1. **Singular**
    1. **Optional (_RECOMMENDED_)**: Optional field normally has pointer to its field which will help you detect if value was actually set or not. Default value of field is ignore if it wasn't set.

    2. **Implicit (_NOT RECOMMENDED_)**: Implicit field contains no pointer play. If any non-zero field is parsed from wire or set, it'll be treated normally. BUT if default value is parsed from wire, it's treated as not set and gets ommitted. 
2. **Repeated**  
This is like slice in go, arraylist in java. Default value of this is empty slice. 
- In proto3, repeated fields of scalar numeric types use packed encoding by default. What's packed encoding? See here [Encoding](#how-protocol-buffers-encode-data)
3. **Maps**  
Map is like key value pair. See here [Maps](#maps)

- Adding `optional` with Message Field doesn't change anything. Here, adding `optional` with Message1 won't make any difference in parsing. Mesage2 and Message3 will be parsed in same way.
```protobuf
message Message1 {}

message Message2 {
  Message1 foo = 1;
}

message Message3 {
  optional Message1 bar = 1;
}
```

## Enums
Link: https://protobuf.dev/programming-guides/proto3/#enum  
- First item in any enum must have 0 value. 
- First item should have name _*_UNSPECIFIED_ or _*_UNNOWN_
- You can define aliases for same enum. This can be done by assigning same values to different enums. Here `MP` and `INDORE` are aliases for each other.
- To use aliases, you should set `allow_alias = true;` in enum.
```protobuf
enum DESTINATION {
    option allow_alias = true;
    DESTINATION_UNSPECIFIED = 0;
    DESTINATION_VARANASI = 1;
    DESTINATION_MP = 2;
    DESTINATION_INDORE = 2;
    DESTINATION_HYD = 3;
}
```
- During serialization, any of the alias can be used. During Deserialization, only first enum is considered.

```
"ENAA_RUNNING" uses the same enum value as "ENAA_STARTED". If this is intended, set 'option allow_alias = true;' to the enum definition. The next available enum value is 3.
```

## How to move .proto files effectively
By default, you can use definitions only from directly imported .proto files. However, sometimes you may need to move a .proto file to a new location. Instead of moving the .proto file directly and updating all the call sites in a single change, you can put a placeholder .proto file in the old location to forward all the imports to the new location using the import public notion.

One thing that's not metioned in document but should be is that if you are importing the *.pb.go, you should have moved enum different name than old one. 

Also, `proto/generated_code/enum_demo.pb.go` is able to resolve `proto/generated_code/enum/enum_demo.pb.go` but `proto/generated_code/addressbook.pb.go` isn't able to resolve same `proto/generated_code/enum/enum_demo.pb.go`

So for now, I've to change the file name of `proto/enums/enum_demo.proto` to unique name `proto/enums/new_enum_demo.proto`. Result: Doesn't work. Probably it's about directory structure.
Now, I've set `option go_package = "proto/generated_code";` for all the proto. It's working fine. Need to look into how `import public` works.

## Nested messages
You can nest messages inside. For Example, you have `PhoneNumber` message inside `Person` message. If you want to access it in any other proto, you can use it as `Person.PhoneNumber`. 
```protobuf
message PhoneBook {
  repeated Person.PhoneNumber phone_numbers = 1;
}
```

## OneOf usage
You may encouter a scenario where you have multiple fields but under any circumstance, it's guaranteed that only one of these fields will be set. Example: If there is field for aadhar number, in case aadhar is not yet done, you have option of setting registration but there is not point of keeping both. 

- Fields in oneofs must not have labels (required/optional/repeated)
- If you set any field, previously set field will be reset (if any).
- You cannot have `map` or `repeated` field inside Oneof. 
- If you wish to have `repeated` field, have that in another message and use in `oneof` field. 
- Be mindset of removing `oneof` fields as it can introduce backward compatibility issue.
- You can use below snippet to get which field is set.
```protobuf
aadharOrRegistrationDescriptor := demo.ProtoReflect().Descriptor().Oneofs().ByName("aadhar_or_registration")
field := demo.ProtoReflect().WhichOneof(aadharOrRegistrationDescriptor)
```

## Defining services
There isn't much detail on proto. Will add details from gRPC. For starting, this is how you can define a RPC service.
```protobuf
service PersonService {
    rpc Search(SearchRequest) returns (SearchResponse);
}
```

## Reading list
- [ ] [Explains what's generated in go protobuf compliation](https://protobuf.dev/reference/go/go-generated/)  
- [ ] [Message wire format](https://protobuf.dev/programming-guides/encoding/)
- [ ] [gRPC](#https://grpc.io/)