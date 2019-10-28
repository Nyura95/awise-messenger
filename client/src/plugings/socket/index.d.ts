declare class AwiseSocket {
  constructor(uri: string, logger?: boolean);
  public send(message: string): void;
  public close(): void;
  public initConversation(token: string, callback: Function): void;
  public onclose: ((this: WebSocket, ev: CloseEvent) => any) | null;
  public onerror: ((this: WebSocket, ev: Event) => any) | null;
  public message: ((message: string) => any) | null;
  public connection: ((user: string) => any) | null;
  public disconnection: ((user: string) => any) | null;
  public error: ((lockey: string, message: string) => any) | null;
}
export = AwiseSocket;
