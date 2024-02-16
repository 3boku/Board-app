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
  
  
    function goToEditPage(id) {
      window.location.href = `/edit/${id}`; // 수정 페이지로 이동
    }
  </script>
  
  <h2>게시물 목록</h2>
  <ul>
    {#each posts as post (post.id)}
      <li>
        <h3>{post.title}</h3>
        <p>{post.content}</p>
        <button on:click={() => goToEditPage(post.id)}>
            <a href={`/update/${post.id}`}>수정</a>
        </button>
      </li>
    {/each}
  </ul>
  