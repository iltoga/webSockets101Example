import { useEffect, useState } from 'react'
import logo from './react-alt.svg'
import './App.css'
 
import * as WebSocket from "websocket"
 
 
 
function App() {
  const [socket, setSocket] = useState<WebSocket.w3cwebsocket | null>(null);
  const [message, setMessage] = useState("");
  const [response, setResponse] = useState("");
  let clientID: string
    

  useEffect(() => {
    clientID = rndStr()
    console.log("in it")
    const socket = new WebSocket.w3cwebsocket('ws://localhost:8080/ws');
    socket.onopen = function () {
      socket.send(`Hi from Client ${clientID}!`)
      socket.onmessage = (msg: any) => {
        console.log(msg);
        console.log("we got msg..")
        setResponse(msg.data);
      };
    };
    setSocket(socket);

    socket.onmessage = (e) => {
        setMessage("Get message from server: " + e.data)
    };
  
    return () => {
        socket.close()
    }   
  }, []); 

  const rndStr = () => {
    return Math.floor((1 + Math.random()) * 0x10000)
        .toString(16)
        .substring(1);
  }

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    if (socket) {
      socket.send(message);
    }
  }

  const handleMessageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setMessage(event.target.value);
  }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>Hello Golang + React Websocket 101 app!</p>
        <form onSubmit={handleSubmit}>
          <input type="text" value={message} onChange={handleMessageChange} />
          <button type="submit">Send</button>
        </form>
        <p>Server response: {response}</p>
        <p>
          Edit <code>App.tsx</code> and save to test HMR updates.
        </p>
        <p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
          {' | '}
          <a
            className="App-link"
            href="https://github.com/iltoga"
            target="_blank"
            rel="noopener noreferrer"
          >
            iltoga on github
          </a>
        </p>
      </header>
    </div>
  )
}

export default App
