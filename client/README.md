# Messenger GoLang

```tsx
import { AwiseSocket } from './awiseSocket';

const Example: IHook<RouteComponentProps> = () => {
  const [messages, setMessages] = React.useState<IMessage[]>([]);
  const [message, setMessage] = React.useState<string>('');
  const [token, setToken] = React.useState<string>('');

  const awiseSocket = React.useMemo(() => {
    return new AwiseSocket('ws://127.0.0.1:3001');
  }, []);

  const init = React.useCallback(() => {
    awiseSocket.init().then(() => {
      awiseSocket.sendMessage('onload', JSON.stringify({ token })); // init user with token
    });
  }, [token]);

  React.useEffect(() => {
    fetch<IConversation>('/v1/conversation/1').then(result => {
      setMessages(result.Data.Messages); // retriveal messages
    });
    return () => {
      awiseSocket.close(); // close when the hooks is destroy
    };
  }, []);

  React.useEffect(() => {
    awiseSocket.onmessage = data => {
      // Action send server side
      switch (
        data.Action // switch action
      ) {
        case 'newMessage':
          setMessages([...messages, data.Data as IMessage]); // add message receive
          break;
      }
    };
  }, [messages]);

  const sendMessage = React.useCallback(
    (message: string) => () => {
      awiseSocket.sendMessage('send', JSON.stringify({ message })); // send message
      setMessage(''); // reset input
    },
    []
  );

  return (
    <Row className={styles.container}>
      <Col lg="12" className={styles.container_button}>
        <Input type="text" value={token} onChange={token => setToken(token)} />
        <Button.Rectangle onClick={init}>Envoyer</Button.Rectangle>
      </Col>
      {messages.map((message, key) => (
        <Col key={key} lg="12" className={styles.container_button}>
          {message.Message}
        </Col>
      ))}
      <Col lg="12" className={styles.container_button}>
        <Input
          type="text"
          value={message}
          onChange={message => setMessage(message)}
        />
        <Button.Rectangle onClick={sendMessage(message)}>
          Envoyer
        </Button.Rectangle>
      </Col>
    </Row>
  );
};

Example.defaultProps = {};

export default Example;
```
