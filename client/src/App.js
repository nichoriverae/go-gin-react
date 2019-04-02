import React, { Component } from 'react';
import Blog from './views/Blog.jsx';
import logo from './logo.svg';
import './App.css';
import { BlogApi } from './api/blog';
import fetchRequest from './utils';
import { apiString } from './utils/config';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      blogposts: [],
      isLoaded: false
    };
  }

  componentDidMount() {
    const blogApi = new BlogApi({ fetch: fetchRequest, apiString: apiString });
    blogApi
      .getAll()
      .then(blogposts => this.setState({ blogposts: blogposts[0] }))
      .catch(error => console.log(error));
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
        </header>
        <section className="body__content columns">
          <div className="column"></div>
          <div className="column">
            <h1>Here are some blog posts</h1>
            <Blog blogposts={this.state.blogposts} />
          </div>
          <div className="column"></div>
        </section>
      </div>
    );
  }
}

export default App;
