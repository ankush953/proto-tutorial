# proto-tutorial
## Table of contents
[Best Practices](#best-pratices)  
[Default Field Values](#default-fields)  
[Updating a mesage type](#updating-a-message-type)  
[Extending a Protocol Buffer](#extending-a-protocol-buffer)  
[Exceptions for protocol buffer extension](#exceptions)

## Best Pratices
Link: https://protobuf.dev/best-practices/dos-donts/   

## Updating a Message Type
Link: https://protobuf.dev/programming-guides/proto3/#updating  
You will need to update the proto message defintions from time to time based on requirements. In this case, Please read below in order
1. [Best Practices](#best-pratices) 
2. [Updating Mesage Type](#updating-a-message-type)
3. [Extending protocol buffer](#extending-a-protocol-buffer)
4. [Exceptions for extension rules](#exceptions)

## Default Fields
Link: https://protobuf.dev/programming-guides/proto3/#default  
- For strings, the default value is the empty string.
- For bytes, the default value is empty bytes.
- For bools, the default value is false.
- For numeric types, the default value is zero.
- For message fields, the field is not set. Its exact value is language-dependent. See the generated code guide for details.
- For enums, the default value is the first defined enum value, which must be 0.

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
