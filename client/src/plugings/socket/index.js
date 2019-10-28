/**
 * Awise socket client
 * @param {string} uri
 * @version 1.0.0
 * @example AwiseSocket('wss://messenger.awise.co')
 * @author Nyura95
 *
 */
function AwiseSocket(uri, logger = true) {
  if (typeof uri !== 'string' && (uri.indexOf('ws://') === -1 || uri.indexOf('wss://') === -1)) {
    throw 'This uri is not correct, you must pass a ws or wss uniform resource identifier';
  }
  Object.defineProperty(this, '_uri', {
    value: uri,
    writable: false,
  });
  this._tokenConversation = null;
  this.webSocket = null;
  this.logger = logger;
  this.onerror = function() {};
  this.onclose = function() {};

  // private action
  this.message = function() {};
  this.connection = function() {};
  this.disconnection = function() {};
  this.error = function() {};
}

/**
 * Initialize a new connexion with a conversation
 * @param {string} token
 * @param {function} callback
 * @version 1.0.0
 * @returns {void}
 * @author Nyura95
 */
AwiseSocket.prototype.initConversation = function(token, callback) {
  if (this.webSocket && this.webSocket.readyState === WebSocket.OPEN) {
    this.close();
  }
  this._tokenConversation = token;
  this._log('init conversation');
  this.webSocket = new WebSocket(this._uri + '/' + this._tokenConversation);
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
    var messages = event.data.split('\n');
    for (let i = 0; i < messages.length; i++) {
      var message = JSON.parse(messages[i]);
      this._log('new message receive (' + message.Action + ')');
      if (message.Action === 'Message') {
        this.message ? this.message(message.Message) : null;
      }
      if (message.Action === 'Connection') {
        this.connection ? this.connection(message.User) : null;
      }
      if (message.Action === 'Disconnection') {
        this.disconnection ? this.disconnection(message.User) : null;
      }
      if (message.Action === 'Error') {
        this.error ? this.error(message.LocKey, message.Message) : null;
      }
    }
  }.bind(this);
};

/**
 *
 * Send a message to the conversation
 * @param {string} token
 * @version 1.0.0
 * @returns {void}
 * @author Nyura95
 */
AwiseSocket.prototype.send = function(message) {
  if (!this._tokenConversation && typeof this._tokenConversation !== 'number') {
    console.warn('You must target a conversation (use initConversation)');
    return;
  }
  this._log('send message : (' + message + ')');
  this.webSocket ? this.webSocket.send(message) : null;
};

/**
 * Close the server connexion
 * @version 1.0.0
 * @returns {void}
 * @author Nyura95
 */
AwiseSocket.prototype.close = function() {
  this._log('close');
  this.webSocket.close(1000, 'close user');
  this.webSocket = null;
};

AwiseSocket.prototype._log = function(...messages) {
  if (console) {
    console.log(`[${this._tokenConversation || 'noToken'}]:`, ...messages);
  }
};

module.exports = AwiseSocket;

/// <reference path=”./index.d.ts” />
