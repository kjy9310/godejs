import React from 'react';
import logo from './logo.svg';
import './App.css';

function App() {

    let loc = window.location;
    let uri = 'ws:';

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
      const out = document.getElementById('output');
      out.innerHTML += evt.data + '<br>';
    }

    setInterval(function() {
	    ws.send(Math.floor(new Date().getTime()/1000) + 'Hello, Server!');
    }, 1000);

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
    </div>
  );
}

export default App;
