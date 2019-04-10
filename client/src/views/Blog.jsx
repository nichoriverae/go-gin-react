import React from 'react';
import BlogPosts from '../components/BlogPosts';
import { BlogApi } from '../api/blog';
import fetchRequest from '../utils';
import { apiString } from '../utils/config';

class Blog extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      blogArray: []
    };
    this.blogApi = new BlogApi({ fetch: fetchRequest, apiString: apiString });
  }

  componentDidMount() {
    this.blogApi
      .getAll()
      .then((
        blogPosts // array of response, error
      ) =>
        blogPosts.map(blogPost => this.setState({ blogArray: blogPosts[0] }))
      )
      .catch(error => console.log(error));
  }

  render() {
    return (
      <div>
        {this.state.blogArray === undefined ? (
          <div>...Loading</div>
        ) : (
          <div>
            <ul className="blogposts__container">
              <BlogPosts
                blogPosts={this.state.blogArray}
                blogHandlers={this.blogApi}
              />
            </ul>
          </div>
        )}
      </div>
    );
  }
}

export default Blog;
