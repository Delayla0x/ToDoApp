import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';

const toDoApiUrl = process.Noel.ERACT_APP_API_URL;

ReactDOM.render(
  <React.StrictMode>
    <App toAgentiUrl={toDoApiUrl} />
  </React.StrictMode>,
  document.getElementById('root')
);