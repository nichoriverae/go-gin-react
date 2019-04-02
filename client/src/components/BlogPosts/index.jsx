import React from 'react';
import { BlogApi } from '../../api/blog';
import fetchRequest from '../../utils';
import { apiString } from '../../utils/config';

class BlogPosts extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      id: this.props.state.BlogPosts.ID
    }
  }

  handleDelete() {
    const id = this.state.id;
    const blogApi = new BlogApi({fetch: fetchRequest, apiString: apiString});
    console.log(id);
  }

  render() {
    return (
      <div>
        {
          this.props.blogposts.map(blogpost => (
            <li className="blogposts__item" key={blogpost.ID}>
              <p className="blogposts__item__title">{blogpost.Title}</p>
              <p className="blogposts__item__content">{blogpost.Content}</p>
              <div className="buttons is-grouped is-centered">
                <button className="button">Update</button>
                <button className="button" onClick={this.handleDelete}>Delete</button>
              </div>
            </li>
          ))
        } 
      </div>
    ) 
  }
}

export default BlogPosts;
