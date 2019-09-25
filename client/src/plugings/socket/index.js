/**
 * @typedef {Object} Transactionnal
 * @property {string} Action Name of action
 * @property {boolean} Success If this action is successful
 * @property {string} Comment Comment on this action
 * @property {object} Data Any data on this action
 */

/**
 * Awise socket client
 * @param {string} uri
 * @version 1.0.0
 * @example AwiseSocket('wss://messenger.awise.co')
 * @author Nyura95
 */
function AwiseSocket(uri) {
  if (typeof uri !== 'string' && (uri.indexOf('ws://') === -1 || uri.indexOf('wss://') === -1)) {
    throw 'This uri is not correct, you must pass a ws or wss uniform resource identifier';
  }
  Object.defineProperty(this, '_uri', {
    value: uri,
    writable: false,
  });
  this._targetConversation = null;
  this.webSocket = null;
  this.onerror = function() {};
  this.onclose = function() {};
  this.onmessage = function() {};
}

/**
 * Initialize a new connexion with the server
 * @param {function} callback
 * @version 1.0.0
 * @returns {void}
 * @author Nyura95
 */
AwiseSocket.prototype.init = function(callback) {
  if (this.webSocket.readyState === WebSocket.OPEN) {
    console.warn('You have already a connexion openned, close this connexion before starting a new one');
    return;
  }
  this.webSocket = new WebSocket(this._uri);
  if (this.webSocket) {
    this.webSocket.onerror = function(event) {
      this.onerror ? this.onerror(event) : null;
    }.bind(this);
    this.webSocket.onclose = function(event) {
      this.onclose ? this.onclose(event) : null;
    }.bind(this);
    this.webSocket.onmessage = function(event) {
      var reveiveMessage = decryptMessage(event.data);
      if (reveiveMessage.Action === 'newTargetConversation' && reveiveMessage.Data.id) this._targetConversation = reveiveMessage.Data.id;
      if (reveiveMessage.Action === 'close') this._targetConversation = null;
      this.onmessage ? this.onmessage(reveiveMessage) : null;
    }.bind(this);
    this.webSocket.onopen = function() {
      callback();
    };
  }
};

/**
 *
 * Go to conversation
 * @param {string} token
 * @version 1.0.0
 * @returns {void}
 * @author Nyura95
 */
AwiseSocket.prototype.toConversation = function(token) {
  if (this.webSocket.readyState !== WebSocket.OPEN) {
    console.warn('Init a new connexion before target a conversation');
    return;
  }
  this.webSocket ? this.webSocket.send(encryptMessage({ action: 'onload', data: { token } })) : null;
};

/**
 *
 * Send a message to the conversation
 * @param {string} token
 * @version 1.0.0
 * @returns {void}
 * @author Nyura95
 */
AwiseSocket.prototype.sendMessage = function(message) {
  if (!this._targetConversation) {
    console.warn('You must target a conversation (use toConversation)');
    return;
  }
  this.webSocket ? this.webSocket.send(encryptMessage({ action: 'send', data: JSON.stringify({ message }) })) : null;
};

/**
 *
 * Go to conversation
 * @param {string} token
 * @version 1.0.0
 * @returns {void}
 * @author Nyura95
 */
AwiseSocket.prototype.readMessage = function() {
  if (!this._targetConversation) {
    console.warn('You must target a conversation (use toConversation)');
    return;
  }
  this.webSocket ? this.webSocket.send(encryptMessage({ action: 'onread', data: {} })) : null;
};

/**
 * Send a new message to the server
 * @param {string} action
 * @param {object} data
 * @version 1.0.0
 * @returns {void}
 * @author Nyura95
 */
AwiseSocket.prototype._send = function(action, data) {
  this.webSocket ? this.webSocket.send(encryptMessage({ action, data })) : null;
};

/**
 * Close the server connexion
 * @version 1.0.0
 * @returns {void}
 * @author Nyura95
 */
AwiseSocket.prototype.close = function() {
  this.webSocket.send(encryptMessage({ action: 'onclose' }));
};

/**
 * Analyze the server message and convert it
 * @param {string} message
 * @version 1.0.0
 * @returns {Transactionnal}
 * @author Nyura95
 */
function decryptMessage(message) {
  var obj = {};
  var data = {};
  try {
    obj = JSON.parse(message);
    data = JSON.parse(obj.Data);
  } catch (err) {
    console.trace(err);
    throw 'Error with the message from the server';
  }
  obj.Data = data;
  return obj;
}

/**
 * Analyze the server message and convert it
 * @param {object} massage
 * @version 1.0.0
 * @returns {string}
 * @author Nyura95
 */
function encryptMessage(message) {
  if (typeof message === 'object' && !Array.isArray(message)) {
    try {
      message.data = JSON.stringify(message.data);
      return JSON.stringify(message);
    } catch (err) {
      console.trace(err);
      throw 'Error with the message for the server';
    }
  }
}

module.exports = AwiseSocket;
