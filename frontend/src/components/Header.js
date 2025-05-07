import React from 'react';
import { Layout } from 'antd';
import { useNavigate } from 'react-router-dom';
const { Header } = Layout;

const HeaderCLick = () => {
    const navigate = useNavigate();

    return (
      <Header
        style={{
          color: '#fff',
          fontSize: '20px',
          cursor: 'pointer',
          display: 'flex',
          justifyContent: 'space-between',
          alignItems: 'center'
        }}
        onClick={() => navigate('/')}
      >
        <span>📝 Article Admin</span>
      </Header>
    );
  };

  export default HeaderCLick;