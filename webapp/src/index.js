import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Header from './components/Header/Header'
import * as serviceWorker from './serviceWorker';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';

import AddProject from './components/AddProject/AddProject'

ReactDOM.render(
  <React.StrictMode>
    <div>
      <Header />
      <div>
        <Router>
          <Switch>
            <Route exact path='/addProject' component={AddProject} />
          </Switch>
        </Router>
      </div>
    </div>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
