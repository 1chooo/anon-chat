import React, { Component } from "react";
import "./ChatInput.scss";

class ChatInput extends Component {
    handleKeyDown = (e) => {
        // 檢查按下的鍵是否是 Enter
        if (e.key === 'Enter') {
            // 調用 send 函數
            this.props.send();
        }
    };

    render() {
        return (
            <div className="ChatInput">
                <input onKeyDown={this.handleKeyDown} />
            </div>
        );
    }
}

export default ChatInput;