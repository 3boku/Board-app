<script>
    import { onMount } from 'svelte';
    import { page } from '@sveltejs/kit';
    let { params } = $page;
    let id = params.id;
    let title = '';
    let content = '';
    
    onMount(async () => {
      const response = await fetch(`http://localhost:8080/posts/${id}`);
      const post = await response.json();
      title = post.title;
      content = post.content;
    });
  
    async function submitPost() {
      const response = await fetch(`http://localhost:8080/update/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ title, content })
      });
  
      if (response.ok) {
        title = '';
        content = '';
        alert('성공적으로 수정되었습니다.');
      } else {
        alert('오류가 발생했습니다.');
      }
    }
  </script>
  
  <h2>글 수정하기</h2>
  <form on:submit|preventDefault={submitPost}>
    <label>
      제목: <input type="text" bind:value={title} required>
    </label>
    <label>
      내용: <textarea bind:value={content} required></textarea>
    </label>
    <button type="submit">수정하기</button>
  </form>
  