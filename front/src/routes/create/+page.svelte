<script>
    let title = '';
    let content = '';
  
    async function submitPost() {
      const response = await fetch('http://localhost:8080/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ title, content })
      });
  
      if (response.ok) {
        title = '';
        content = '';
        alert('성공적으로 게시되었습니다.');
      } else {
        alert('오류가 발생했습니다.');
      }
    }
  </script>
  
  <h2>글쓰기</h2>
  <form on:submit|preventDefault={submitPost}>
    <label>
      제목: <input type="text" bind:value={title} required>
    </label>
    <label>
      내용: <textarea bind:value={content} required></textarea>
    </label>
    <button type="submit">작성하기</button>
  </form>
  
  
  <style>
    h2 {
        text-align: center;
    }
    form {
        display: flex;
        justify-content: center;
        flex-direction: co;
    }
  </style>