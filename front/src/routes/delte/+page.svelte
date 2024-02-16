<script>
    import { onMount } from 'svelte';
    let posts = []; // 초기값을 빈 배열로 설정
  
    onMount(async () => {
      await loadPosts();
    });
  
    async function loadPosts() {
      const response = await fetch('http://localhost:8080/posts');
      posts = await response.json();
    }
  
    async function deletePost(id) {
      await fetch(`http://localhost:8080/delete/${id}`, {
        method: 'DELETE',
      });
      await loadPosts(); // 게시물 삭제 후 게시물 목록을 다시 불러옵니다.
    }
  </script>
  
  <h2>게시물 목록</h2>
  <ul>
    {#each posts as post (post.id)}
      <li>
        <h3>{post.title}</h3>
        <p>{post.content}</p>
        <button on:click={() => deletePost(post.id)}>삭제</button>
      </li>
    {/each}
  </ul>
  