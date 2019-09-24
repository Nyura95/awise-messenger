function AwiseSocket(uri) {
  this.uri = uri;
  this.webSocket = null;
  this.onerror = function() {};
  this.onclose = function() {};
  this.onmessage = function() {};
}

AwiseSocket.prototype.init = function(callback) {
  this.webSocket = new WebSocket(this.uri);
  if (this.webSocket) {
    this.webSocket.onopen = function() {
      callback();
    };
    this.webSocket.onerror = function(event) {
      this.onerror ? this.onerror(event) : null;
    }.bind(this);
    this.webSocket.onclose = function(event) {
      this.onclose ? this.onclose(event) : null;
    }.bind(this);
    this.webSocket.onmessage = function(event) {
      const reveiveMessage = this._receiveMessage(event.data);
      this.onmessage ? this.onmessage(reveiveMessage) : null;
    }.bind(this);
  }
};

AwiseSocket.prototype.sendMessage = function(action, data) {
  this.webSocket ? this.webSocket.send(JSON.stringify({ action, data })) : null;
};

AwiseSocket.prototype.close = function() {
  this.webSocket ? this.webSocket.close() : null;
};

AwiseSocket.prototype._receiveMessage = function(receiveMessage) {
  const obj = JSON.parse(receiveMessage);
  return {
    ...obj,
    Data: JSON.parse(obj.Data),
  };
};

module.exports = AwiseSocket;
