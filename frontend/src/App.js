import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { Layout } from 'antd';
import AllPosts from './components/AllPosts';
import AddNew from './components/AddNew';
import Preview from './components/Preview';
import ArticleForm from './components/ArticleForm';
import HeaderCLick from './components/Header';

const { Header, Content, Footer } = Layout;

function App() {
  return (
    <Router>
      <Layout>
        <HeaderCLick style={{ color: '#fff', fontSize: '20px' }}>Article Admin</HeaderCLick>
        <Content style={{ padding: '20px' }}>
          <Routes>
            <Route path="/all-posts" element={<AllPosts />} />
            <Route path="/add-new" element={<AddNew />} />
            <Route path="/preview" element={<Preview />} />
            <Route path="/edit/:id" element={<ArticleForm />} />
            <Route path="/" element={<AllPosts />} />
          </Routes>
        </Content>
        <Footer>© 2025 Article System</Footer>
      </Layout>
    </Router>
  );
}

export default App;