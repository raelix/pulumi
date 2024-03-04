# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pulumi/testing/language.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dpulumi/testing/language.proto\x12\x11pulumirpc.testing\"\x19\n\x17GetLanguageTestsRequest\")\n\x18GetLanguageTestsResponse\x12\r\n\x05tests\x18\x01 \x03(\t\"\xb0\x01\n\x1bPrepareLanguageTestsRequest\x12\x1c\n\x14language_plugin_name\x18\x01 \x01(\t\x12\x1e\n\x16language_plugin_target\x18\x02 \x01(\t\x12\x1a\n\x12snapshot_directory\x18\x03 \x01(\t\x12\x1b\n\x13temporary_directory\x18\x04 \x01(\t\x12\x1a\n\x12\x63ore_sdk_directory\x18\x05 \x01(\t\"-\n\x1cPrepareLanguageTestsResponse\x12\r\n\x05token\x18\x01 \x01(\t\"5\n\x16RunLanguageTestRequest\x12\r\n\x05token\x18\x01 \x01(\t\x12\x0c\n\x04test\x18\x02 \x01(\t\"\\\n\x17RunLanguageTestResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x12\x10\n\x08messages\x18\x02 \x03(\t\x12\x0e\n\x06stdout\x18\x03 \x01(\t\x12\x0e\n\x06stderr\x18\x04 \x01(\t2\xe4\x02\n\x0cLanguageTest\x12m\n\x10GetLanguageTests\x12*.pulumirpc.testing.GetLanguageTestsRequest\x1a+.pulumirpc.testing.GetLanguageTestsResponse\"\x00\x12y\n\x14PrepareLanguageTests\x12..pulumirpc.testing.PrepareLanguageTestsRequest\x1a/.pulumirpc.testing.PrepareLanguageTestsResponse\"\x00\x12j\n\x0fRunLanguageTest\x12).pulumirpc.testing.RunLanguageTestRequest\x1a*.pulumirpc.testing.RunLanguageTestResponse\"\x00\x42\x32Z0github.com/pulumi/pulumi/sdk/v3/proto/go/testingb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'pulumi.testing.language_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z0github.com/pulumi/pulumi/sdk/v3/proto/go/testing'
  _GETLANGUAGETESTSREQUEST._serialized_start=52
  _GETLANGUAGETESTSREQUEST._serialized_end=77
  _GETLANGUAGETESTSRESPONSE._serialized_start=79
  _GETLANGUAGETESTSRESPONSE._serialized_end=120
  _PREPARELANGUAGETESTSREQUEST._serialized_start=123
  _PREPARELANGUAGETESTSREQUEST._serialized_end=299
  _PREPARELANGUAGETESTSRESPONSE._serialized_start=301
  _PREPARELANGUAGETESTSRESPONSE._serialized_end=346
  _RUNLANGUAGETESTREQUEST._serialized_start=348
  _RUNLANGUAGETESTREQUEST._serialized_end=401
  _RUNLANGUAGETESTRESPONSE._serialized_start=403
  _RUNLANGUAGETESTRESPONSE._serialized_end=495
  _LANGUAGETEST._serialized_start=498
  _LANGUAGETEST._serialized_end=854
# @@protoc_insertion_point(module_scope)
