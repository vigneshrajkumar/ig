import React from "react"


function CommentInput() {
    return (
        <div className="comment-input">  
            <div> <input type="text" placeholder="what do you think?" /> </div>
            <div> <button> Send </button> </div>
        </div>
    )
}

export default CommentInput