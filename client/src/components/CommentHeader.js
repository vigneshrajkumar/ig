import React from 'react'

function CommentHeader(){
    return (
        <div className="comment-header">
            <div> <img src="https://via.placeholder.com/48" alt="user-thumbnail"></img> </div>
            <div className="comment-header-name-container">
                <div>Ann Hathway</div>
                <div>15 minutes ago</div>
            </div>
            <div>18 Hearts</div>
        </div>
    )
}

export default CommentHeader