import React from 'react';
import View from './View';
import Proptypes from 'prop-types';

class BlogPosts extends React.Component {
  static propTypes = {
    blogHandlers: Proptypes.object.isRequired,
    blogPosts: Proptypes.array.isRequired
  }

  handleChange = id => {
    this.props.blogHandlers.createOrUpdate(id);
  };

  handleDelete = id => {
    this.props.blogHandlers.delete(id);
  };

  render() {
    const {blogPosts} = this.props;

    return (
      <div>
        <View
          blogPosts={blogPosts}
          handleChange={this.handleChange}
          handleDelete={this.handleDelete}
        />
      </div>
    );
  }
}

export default BlogPosts;
