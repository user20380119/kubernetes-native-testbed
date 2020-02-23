// source: comment.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
goog.object.extend(proto, google_protobuf_empty_pb);
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.commentpb.Comment', null, global);
goog.exportSymbol('proto.commentpb.DeleteRequest', null, global);
goog.exportSymbol('proto.commentpb.GetRequest', null, global);
goog.exportSymbol('proto.commentpb.GetResponse', null, global);
goog.exportSymbol('proto.commentpb.SetRequest', null, global);
goog.exportSymbol('proto.commentpb.SetResponse', null, global);
goog.exportSymbol('proto.commentpb.UpdateRequest', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.commentpb.Comment = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.commentpb.Comment, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.commentpb.Comment.displayName = 'proto.commentpb.Comment';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.commentpb.GetRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.commentpb.GetRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.commentpb.GetRequest.displayName = 'proto.commentpb.GetRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.commentpb.GetResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.commentpb.GetResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.commentpb.GetResponse.displayName = 'proto.commentpb.GetResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.commentpb.SetRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.commentpb.SetRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.commentpb.SetRequest.displayName = 'proto.commentpb.SetRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.commentpb.SetResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.commentpb.SetResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.commentpb.SetResponse.displayName = 'proto.commentpb.SetResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.commentpb.UpdateRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.commentpb.UpdateRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.commentpb.UpdateRequest.displayName = 'proto.commentpb.UpdateRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.commentpb.DeleteRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.commentpb.DeleteRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.commentpb.DeleteRequest.displayName = 'proto.commentpb.DeleteRequest';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.commentpb.Comment.prototype.toObject = function(opt_includeInstance) {
  return proto.commentpb.Comment.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.commentpb.Comment} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.commentpb.Comment.toObject = function(includeInstance, msg) {
  var f, obj = {
    uuid: jspb.Message.getFieldWithDefault(msg, 1, ""),
    message: jspb.Message.getFieldWithDefault(msg, 2, ""),
    useruuid: jspb.Message.getFieldWithDefault(msg, 3, ""),
    parentcommentuuid: jspb.Message.getFieldWithDefault(msg, 4, ""),
    createdat: (f = msg.getCreatedat()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    updatedat: (f = msg.getUpdatedat()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    deletedat: (f = msg.getDeletedat()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.commentpb.Comment}
 */
proto.commentpb.Comment.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.commentpb.Comment;
  return proto.commentpb.Comment.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.commentpb.Comment} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.commentpb.Comment}
 */
proto.commentpb.Comment.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setUuid(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setMessage(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setUseruuid(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setParentcommentuuid(value);
      break;
    case 5:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setCreatedat(value);
      break;
    case 6:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setUpdatedat(value);
      break;
    case 7:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setDeletedat(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.commentpb.Comment.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.commentpb.Comment.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.commentpb.Comment} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.commentpb.Comment.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUuid();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getMessage();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getUseruuid();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getParentcommentuuid();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getCreatedat();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getUpdatedat();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getDeletedat();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
};


/**
 * optional string UUID = 1;
 * @return {string}
 */
proto.commentpb.Comment.prototype.getUuid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.commentpb.Comment.prototype.setUuid = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string message = 2;
 * @return {string}
 */
proto.commentpb.Comment.prototype.getMessage = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.commentpb.Comment.prototype.setMessage = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string userUUID = 3;
 * @return {string}
 */
proto.commentpb.Comment.prototype.getUseruuid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.commentpb.Comment.prototype.setUseruuid = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string parentCommentUUID = 4;
 * @return {string}
 */
proto.commentpb.Comment.prototype.getParentcommentuuid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.commentpb.Comment.prototype.setParentcommentuuid = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional google.protobuf.Timestamp createdAt = 5;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.commentpb.Comment.prototype.getCreatedat = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 5));
};


/** @param {?proto.google.protobuf.Timestamp|undefined} value */
proto.commentpb.Comment.prototype.setCreatedat = function(value) {
  jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.commentpb.Comment.prototype.clearCreatedat = function() {
  this.setCreatedat(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.commentpb.Comment.prototype.hasCreatedat = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional google.protobuf.Timestamp updatedAt = 6;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.commentpb.Comment.prototype.getUpdatedat = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 6));
};


/** @param {?proto.google.protobuf.Timestamp|undefined} value */
proto.commentpb.Comment.prototype.setUpdatedat = function(value) {
  jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.commentpb.Comment.prototype.clearUpdatedat = function() {
  this.setUpdatedat(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.commentpb.Comment.prototype.hasUpdatedat = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional google.protobuf.Timestamp deletedAt = 7;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.commentpb.Comment.prototype.getDeletedat = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 7));
};


/** @param {?proto.google.protobuf.Timestamp|undefined} value */
proto.commentpb.Comment.prototype.setDeletedat = function(value) {
  jspb.Message.setWrapperField(this, 7, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.commentpb.Comment.prototype.clearDeletedat = function() {
  this.setDeletedat(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.commentpb.Comment.prototype.hasDeletedat = function() {
  return jspb.Message.getField(this, 7) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.commentpb.GetRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.commentpb.GetRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.commentpb.GetRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.commentpb.GetRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    uuid: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.commentpb.GetRequest}
 */
proto.commentpb.GetRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.commentpb.GetRequest;
  return proto.commentpb.GetRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.commentpb.GetRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.commentpb.GetRequest}
 */
proto.commentpb.GetRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setUuid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.commentpb.GetRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.commentpb.GetRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.commentpb.GetRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.commentpb.GetRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUuid();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string UUID = 1;
 * @return {string}
 */
proto.commentpb.GetRequest.prototype.getUuid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.commentpb.GetRequest.prototype.setUuid = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.commentpb.GetResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.commentpb.GetResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.commentpb.GetResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.commentpb.GetResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    comment: (f = msg.getComment()) && proto.commentpb.Comment.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.commentpb.GetResponse}
 */
proto.commentpb.GetResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.commentpb.GetResponse;
  return proto.commentpb.GetResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.commentpb.GetResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.commentpb.GetResponse}
 */
proto.commentpb.GetResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.commentpb.Comment;
      reader.readMessage(value,proto.commentpb.Comment.deserializeBinaryFromReader);
      msg.setComment(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.commentpb.GetResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.commentpb.GetResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.commentpb.GetResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.commentpb.GetResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getComment();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.commentpb.Comment.serializeBinaryToWriter
    );
  }
};


/**
 * optional Comment comment = 1;
 * @return {?proto.commentpb.Comment}
 */
proto.commentpb.GetResponse.prototype.getComment = function() {
  return /** @type{?proto.commentpb.Comment} */ (
    jspb.Message.getWrapperField(this, proto.commentpb.Comment, 1));
};


/** @param {?proto.commentpb.Comment|undefined} value */
proto.commentpb.GetResponse.prototype.setComment = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.commentpb.GetResponse.prototype.clearComment = function() {
  this.setComment(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.commentpb.GetResponse.prototype.hasComment = function() {
  return jspb.Message.getField(this, 1) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.commentpb.SetRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.commentpb.SetRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.commentpb.SetRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.commentpb.SetRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    comment: (f = msg.getComment()) && proto.commentpb.Comment.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.commentpb.SetRequest}
 */
proto.commentpb.SetRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.commentpb.SetRequest;
  return proto.commentpb.SetRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.commentpb.SetRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.commentpb.SetRequest}
 */
proto.commentpb.SetRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.commentpb.Comment;
      reader.readMessage(value,proto.commentpb.Comment.deserializeBinaryFromReader);
      msg.setComment(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.commentpb.SetRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.commentpb.SetRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.commentpb.SetRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.commentpb.SetRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getComment();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.commentpb.Comment.serializeBinaryToWriter
    );
  }
};


/**
 * optional Comment comment = 1;
 * @return {?proto.commentpb.Comment}
 */
proto.commentpb.SetRequest.prototype.getComment = function() {
  return /** @type{?proto.commentpb.Comment} */ (
    jspb.Message.getWrapperField(this, proto.commentpb.Comment, 1));
};


/** @param {?proto.commentpb.Comment|undefined} value */
proto.commentpb.SetRequest.prototype.setComment = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.commentpb.SetRequest.prototype.clearComment = function() {
  this.setComment(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.commentpb.SetRequest.prototype.hasComment = function() {
  return jspb.Message.getField(this, 1) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.commentpb.SetResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.commentpb.SetResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.commentpb.SetResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.commentpb.SetResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    uuid: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.commentpb.SetResponse}
 */
proto.commentpb.SetResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.commentpb.SetResponse;
  return proto.commentpb.SetResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.commentpb.SetResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.commentpb.SetResponse}
 */
proto.commentpb.SetResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setUuid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.commentpb.SetResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.commentpb.SetResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.commentpb.SetResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.commentpb.SetResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUuid();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string UUID = 1;
 * @return {string}
 */
proto.commentpb.SetResponse.prototype.getUuid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.commentpb.SetResponse.prototype.setUuid = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.commentpb.UpdateRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.commentpb.UpdateRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.commentpb.UpdateRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.commentpb.UpdateRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    comment: (f = msg.getComment()) && proto.commentpb.Comment.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.commentpb.UpdateRequest}
 */
proto.commentpb.UpdateRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.commentpb.UpdateRequest;
  return proto.commentpb.UpdateRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.commentpb.UpdateRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.commentpb.UpdateRequest}
 */
proto.commentpb.UpdateRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.commentpb.Comment;
      reader.readMessage(value,proto.commentpb.Comment.deserializeBinaryFromReader);
      msg.setComment(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.commentpb.UpdateRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.commentpb.UpdateRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.commentpb.UpdateRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.commentpb.UpdateRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getComment();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.commentpb.Comment.serializeBinaryToWriter
    );
  }
};


/**
 * optional Comment comment = 1;
 * @return {?proto.commentpb.Comment}
 */
proto.commentpb.UpdateRequest.prototype.getComment = function() {
  return /** @type{?proto.commentpb.Comment} */ (
    jspb.Message.getWrapperField(this, proto.commentpb.Comment, 1));
};


/** @param {?proto.commentpb.Comment|undefined} value */
proto.commentpb.UpdateRequest.prototype.setComment = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.commentpb.UpdateRequest.prototype.clearComment = function() {
  this.setComment(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.commentpb.UpdateRequest.prototype.hasComment = function() {
  return jspb.Message.getField(this, 1) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.commentpb.DeleteRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.commentpb.DeleteRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.commentpb.DeleteRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.commentpb.DeleteRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    uuid: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.commentpb.DeleteRequest}
 */
proto.commentpb.DeleteRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.commentpb.DeleteRequest;
  return proto.commentpb.DeleteRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.commentpb.DeleteRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.commentpb.DeleteRequest}
 */
proto.commentpb.DeleteRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setUuid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.commentpb.DeleteRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.commentpb.DeleteRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.commentpb.DeleteRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.commentpb.DeleteRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUuid();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string UUID = 1;
 * @return {string}
 */
proto.commentpb.DeleteRequest.prototype.getUuid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.commentpb.DeleteRequest.prototype.setUuid = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


goog.object.extend(exports, proto.commentpb);
