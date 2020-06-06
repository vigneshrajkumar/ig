import React from "react"

import CommentInput from "./CommentInput"

import CommentHeader from "./CommentHeader"

import Comment from "./Comment"


class Post extends React.Component {
    render() {
        return (
            <div className="post">
                <div className="post-top">
                    <div className="image-container">
                        <img src="https://via.placeholder.com/300x220" alt="user-thumbnail"></img>
                    </div>
                    <div>
                        <CommentHeader />
                        <div className="comment-list">
                            <Comment />
                            <Comment />
                            <Comment />
                            <Comment />
                            <Comment />
                        </div>
                    </div>
                </div>
                <div className="post-bottom">
                    <div className="caption">
                        Sample caption to test he Rx app
                    </div>
                    <div>
                        <CommentInput />
                    </div>
                </div>
            </div>
        )
    }
}

export default Post