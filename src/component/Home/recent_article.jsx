import React, { Component } from 'react';
import PropTypes from 'prop-types';
// import '../style/recentArticle.sass';

class HomeArticle extends Component {
  constructor(props) {
    super(props);
    this.getProjects = this.getProjects.bind(this);
  }

  getProjects() {
    const { history } = this.props;
    return (
      <div
        className="card projectSummary mt-4"
        onClick={() => { history.push('/project/2'); }}
        >
        <div className="card-body">
          <span className="card-title">Project Title</span>
          <p>Posted Body</p>
          <p className="postDate">3rd Sept 2017, 2:00 am</p>
        </div>
      </div>
    );
  }

  render() {
    return (
      <div className="projectList section">
        <h2>Recent Stories</h2>
        {this.getProjects()}
        {this.getProjects()}
        {this.getProjects()}
      </div>
    );
  }
}

HomeArticle.propTypes = {
  history: PropTypes.shape({
    push: PropTypes.func,
  }).isRequired,
};

export default HomeArticle;
