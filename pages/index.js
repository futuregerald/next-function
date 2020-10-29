import React, { useEffect } from 'react';
import netlifyIdentity from 'netlify-identity-widget';

const Home = () => {
  useEffect(() => {
    netlifyIdentity.init();
  });
  const openSignUp = () => {
    netlifyIdentity.open('login');
  };

  return (
    <div>
      <button type="submit" onClick={openSignUp}>
        Sign-up! please
      </button>

      <style jsx>{`
        :global(html, body) {
          margin: 0;
          padding: 0;
          height: 100%;
        }

        :global(body) {
          font-size: calc(10px + 1vmin);
          font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto',
            'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans',
            'Helvetica Neue', sans-serif;
          -webkit-font-smoothing: antialiased;
          -moz-osx-font-smoothing: grayscale;

          display: flex;
          flex-direction: column;
          align-items: center;
          justify-content: center;
          text-align: center;
          background-color: #282c34;
          color: white;
        }

        a {
          color: pink;
          text-decoration: none;
        }

        .content {
          padding: 0 32px;
        }

        button {
          box-shadow: inset 0px 1px 0px 0px #ffffff;
          background: linear-gradient(to bottom, #f9f9f9 5%, #e9e9e9 100%);
          background-color: #f9f9f9;
          border-radius: 6px;
          border: 1px solid #dcdcdc;
          display: inline-block;
          cursor: pointer;
          color: #666666;
          font-family: Arial;
          font-size: 15px;
          font-weight: bold;
          text-decoration: none;
          text-shadow: 0px 1px 0px #ffffff;
          padding: 1rem;
          padding-left: 10rem;
          padding-right: 10rem;
          font-size: 1.5rem;
        }
      `}</style>
    </div>
  );
};

export default Home;
