import React from 'react';

const PostShow: React.FC = () => {
  return (
    <div className="post-show">
      <h2>投稿一覧</h2>
      <div className="post">
        <h3>投稿タイトル1</h3>
        <p>ここに投稿の内容が表示されます。これはサンプルテキストです。</p>
      </div>
      <div className="post">
        <h3>投稿タイトル2</h3>
        <p>ここに別の投稿の内容が表示されます。これは別のサンプルテキストです。</p>
      </div>
      {/* さらに投稿を追加できます */}
    </div>
  );
};

export default PostShow;
