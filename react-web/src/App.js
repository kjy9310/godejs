import React from 'react';
import logo from './logo.svg';
import './App.css';

function App() {

    let loc = window.location;
    let uri = 'ws:';
	let sentTimestamp, receivedTimestamp = 0;

    if (loc.protocol === 'https:') {
      uri = 'wss:';
    }
    uri += '//' + loc.host;
    uri += loc.pathname + 'ws';

    const ws = new WebSocket(uri)

    ws.onopen = function() {
      console.log('Connected')
    }

    ws.onmessage = function(evt) {
	console.log("time in server : ", new Date(evt.data))
      console.log("server delay : ",parseInt(evt.data)-sentTimestamp);
	receivedTimestamp = Math.floor(new Date().getTime());
	console.log('ping : ',receivedTimestamp-sentTimestamp)
    }

    setInterval(function() {
	sentTimestamp = Math.floor(new Date().getTime());
	    ws.send(Math.floor(new Date().getTime()) + 'Hello, Server!');
    }, 5000);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
      <div id="output"></div>
    </div>
  );
}

export default App;
