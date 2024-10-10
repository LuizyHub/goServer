// static/app.js

document.addEventListener('DOMContentLoaded', function () {
    loadPosts();

    document.getElementById('post-form').addEventListener('submit', function (e) {
        e.preventDefault();
        createPost();
    });
});

function loadPosts() {
    fetch('/api/posts')
        .then(response => response.json())
        .then(data => {
            const postsDiv = document.getElementById('posts');
            postsDiv.innerHTML = '';

            if (data.length === 0) {
                postsDiv.innerHTML = '<p>포스트가 없습니다.</p>';
                return;
            }

            data.forEach(post => {
                const postDiv = document.createElement('div');
                postDiv.className = 'post';

                const title = document.createElement('h2');
                title.textContent = post.title;
                title.onclick = () => viewPost(post.id);

                const content = document.createElement('p');
                content.textContent = post.content.length > 100 ? post.content.substring(0, 100) + '...' : post.content;

                const editButton = document.createElement('button');
                editButton.textContent = '수정';
                editButton.onclick = () => editPost(post.id);

                const deleteButton = document.createElement('button');
                deleteButton.textContent = '삭제';
                deleteButton.onclick = () => deletePost(post.id);

                postDiv.appendChild(title);
                postDiv.appendChild(content);
                postDiv.appendChild(editButton);
                postDiv.appendChild(deleteButton);

                postsDiv.appendChild(postDiv);
            });
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

function createPost() {
    const title = document.getElementById('title').value;
    const content = document.getElementById('content').value;

    fetch('/api/posts', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({title, content})
    })
        .then(response => response.json())
        .then(data => {
            hideCreateForm();
            loadPosts();
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

function viewPost(id) {
    fetch('/api/posts/' + id)
        .then(response => response.json())
        .then(post => {
            alert('제목: ' + post.title + '\n\n내용: ' + post.content);
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

function editPost(id) {
    const newTitle = prompt('새 제목을 입력하세요:');
    if (!newTitle) return;

    const newContent = prompt('새 내용을 입력하세요:');
    if (!newContent) return;

    fetch('/api/posts/' + id, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({title: newTitle, content: newContent})
    })
        .then(() => {
            loadPosts();
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

function deletePost(id) {
    if (!confirm('정말 삭제하시겠습니까?')) return;

    fetch('/api/posts/' + id, {
        method: 'DELETE'
    })
        .then(() => {
            loadPosts();
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

function showCreateForm() {
    document.getElementById('form-container').style.display = 'block';
}

function hideCreateForm() {
    document.getElementById('form-container').style.display = 'none';
    document.getElementById('title').value = '';
    document.getElementById('content').value = '';
}
