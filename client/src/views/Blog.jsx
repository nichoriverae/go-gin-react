import React from 'react';
import BlogPosts from '../components/BlogPosts';
import { BlogApi } from '../api/blog';
import fetchRequest from '../utils';
import { apiString } from '../utils/config';

class Blog extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      title: '',
      content: ''
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    const name = event.target.name;
    const value = event.target.value;

    this.setState({
      [name]: value
    })
  }

  handleSubmit(event) {
    let data = {
      title: this.state.title,
      content: this.state.content
    }
    let blogApi = new BlogApi({ fetch: fetchRequest, apiString: apiString });
    blogApi.create(data).then(response => console.log(response));
    event.preventDefault();
  }

  render() {
    return (
      <div>
        <ul className="blogposts__container">
          <BlogPosts blogposts={this.props.blogposts} />
        </ul>
        <div className="">
          <form onSubmit={this.handleSubmit}>
            <label>
              Title:
              <input type="text" name="title" value={this.state.title} onChange={this.handleChange} />
            </label>
            <label>
              Content:
              <input type="text" name="content" value={this.state.content} onChange={this.handleChange} />
            </label>
            <input type="submit" value="submit"/>
          </form>
        </div>
      </div>
    );
  }
}

export default Blog;
