import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import App from './App';
import PageNotFound from './PageNotFound';

if (process.env.NODE_ENV !== 'production') {
  require('dotenv').config();
}

const Routes = () => (
  <Switch>
    <Route exact path="/" component={App} />
    <Route component={PageNotFound} />
  </Switch>
);

const Root = () => {
  return (
    <Router>
      <Routes />
    </Router>
  );
};

ReactDOM.render(<Root />, document.getElementById('root'));