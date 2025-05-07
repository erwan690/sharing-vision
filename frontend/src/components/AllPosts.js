import React, { useState, useEffect } from 'react';
import { Tabs, Table, Button, Space, message, Card } from 'antd';
import { EditOutlined, DeleteOutlined } from '@ant-design/icons';
import { Link, useNavigate } from 'react-router-dom';
import { getArticlesByStatus, updateArticle, deleteArticle } from '../services/articleService';

const { TabPane } = Tabs;

const AllPosts = () => {
  const navigate = useNavigate();
  const [loading, setLoading] = useState(false);
  const [articles, setArticles] = useState({ Publish: [], Draft: [], Trash: [] });

  const fetchArticles = async () => {
    setLoading(true);
    try {
      const published = await getArticlesByStatus('Publish');
      const drafts = await getArticlesByStatus('Draft');
      const trashed = await getArticlesByStatus('Trash');
      console.log(published, drafts, trashed);
      setArticles({ Publish: published, Draft: drafts, Trash: trashed });
    } catch (error) {
      message.error('Failed to load articles');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchArticles();
  }, []);

  const moveToTrash = async (id) => {
    try {
      await deleteArticle(id);
      message.success('Moved to trash');
      fetchArticles();
    } catch (error) {
      message.error('Failed to move to trash');
    }
  };

  const columns = [
    {
      title: 'Title',
      dataIndex: 'title',
      key: 'title',
    },
    {
      title: 'Category',
      dataIndex: 'category',
      key: 'category',
    },
    {
      title: 'Action',
      key: 'action',
      render: (_, record) => (
        <Space size="middle">
          <Button icon={<EditOutlined />} onClick={() => navigate(`/edit/${record.id}`)} />
          <Button icon={<DeleteOutlined />} onClick={() => moveToTrash(record.id)} danger />
        </Space>
      ),
    },
  ];

  return (
    <Card
      title="All Posts"
      extra={
        <Button type="primary" onClick={() => navigate('/add-new')}>
          Add New
        </Button>
      }
    >
      <Tabs defaultActiveKey="Publish">
        <TabPane tab="Published" key="Publish">
          <Table dataSource={articles.Publish} columns={columns} loading={loading} rowKey="id" />
        </TabPane>
        <TabPane tab="Drafts" key="Draft">
          <Table dataSource={articles.Draft} columns={columns} loading={loading} rowKey="id" />
        </TabPane>
        <TabPane tab="Trashed" key="Trash">
          <Table dataSource={articles.Trash} columns={columns} loading={loading} rowKey="id" />
        </TabPane>
      </Tabs>
    </Card>
  );
};

export default AllPosts;