import React from 'react';

const View = (blogPosts, handleDelete, handleUpdate) => {

  return (
    <div>
      {blogPosts.blogPosts.map(blogPost => (
        <li className="blogposts__item" key={blogPost.ID}>
          <p className="blogposts__item__title">{blogPost.Title}</p>
          <p className="blogposts__item__content">{blogPost.Content}</p>
          <div className="buttons is-grouped is-centered">
            <button className="button" onClick={() => { blogPosts.handleUpdate(blogPost.ID) }}>Update</button>
            <button className="button" onClick={() => { blogPosts.handleDelete(blogPost.ID) }}>
              Delete
            </button>
          </div>
        </li>
      ))}
    </div>
  );
};

export default View;
