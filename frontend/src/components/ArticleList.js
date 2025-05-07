import React, { useState, useEffect } from 'react';
import {
  Card,
  Button,
  Input,
  Space,
  Table,
  Modal,
  Form,
  Input as AntInput,
  Select,
  message
} from 'antd';
import ArticleForm from './ArticleForm';
import apiClient from '../utils/api';


const { Search } = Input;
const { Option } = Select;

const ArticleList = () => {
  const [articles, setArticles] = useState([]);
  const [loading, setLoading] = useState(false);
  const [total, setTotal] = useState(0);
  const [limit, setLimit] = useState(10);
  const [offset, setOffset] = useState(0);
  const [searchTitle, setSearchTitle] = useState('');
  const [searchContent, setSearchContent] = useState('');
  const [searchCategory, setSearchCategory] = useState('');
  const [searchStatus, setSearchStatus] = useState('');
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [currentArticle, setCurrentArticle] = useState(null);

  const fetchArticles = async () => {
    setLoading(true);
    try {
        const rawParams = {
            limit,
            offset,
            title: searchTitle,
            content: searchContent,
            category: searchCategory,
            status: searchStatus
          };

          const params = Object.entries(rawParams).reduce((acc, [key, value]) => {
            if (value !== undefined && value !== null && value !== '') {
              acc[key] = value;
            }
            return acc;
          }, {});

      const response = await apiClient.get('/article', {
        params: params,
      });
      if (response.status) {
        setArticles(response.data || []);
        setTotal(response.data.length);
      } else {
        throw new Error(response.message);
      }
    } catch (error) {
      message.error('Failed to fetch articles');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchArticles();
  }, [limit, offset, searchTitle, searchContent, searchCategory, searchStatus]);

  const handleTableChange = (pagination) => {
    setOffset((pagination.current - 1) * pagination.pageSize);
    setLimit(pagination.pageSize);
  };

  const showCreateModal = () => {
    setCurrentArticle(null);
    setIsModalVisible(true);
  };

  const showEditModal = (article) => {
    setCurrentArticle(article);
    setIsModalVisible(true);
  };

  const handleOk = () => {
    fetchArticles();
    setIsModalVisible(false);
  };

  const handleCancel = () => {
    setIsModalVisible(false);
  };

  const handleDelete = async (id) => {
    try {
      const response = await apiClient.delete(`/article/${id}`);

      if (response.status) {
        message.success('Article deleted successfully');
        fetchArticles();
      } else {
        throw new Error(response.message);
      }
    } catch (error) {
      message.error('Failed to delete article');
      console.error(error);
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
      title: 'Status',
      dataIndex: 'status',
      key: 'status',
    },
    {
      title: 'Created Date',
      dataIndex: 'created_date',
      key: 'created_date',
    },
    {
      title: 'Action',
      key: 'action',
      render: (_, record) => (
        <Space size="middle">
          <Button onClick={() => showEditModal(record)}>Edit</Button>
          <Button danger onClick={() => handleDelete(record.id)}>Delete</Button>
        </Space>
      ),
    },
  ];


  return (
    <Card
      title="Articles"
      extra={
        <Button type="primary" onClick={showCreateModal}>
          Create Article
        </Button>
      }
    >
      <Space direction="vertical" size="large" style={{ width: '100%' }}>
        <Space wrap>
          <Search
            placeholder="Search by title"
            value={searchTitle}
            onChange={(e) => setSearchTitle(e.target.value)}
            style={{ width: 200 }}
          />
          <Search
            placeholder="Search by content"
            value={searchContent}
            onChange={(e) => setSearchContent(e.target.value)}
            style={{ width: 200 }}
          />
          <Search
            placeholder="Search by category"
            value={searchCategory}
            onChange={(e) => setSearchCategory(e.target.value)}
            style={{ width: 200 }}
          />
          <Select
            placeholder="Filter by status"
            value={searchStatus || undefined}
            onChange={(value) => setSearchStatus(value)}
            style={{ width: 200 }}
            allowClear
          >
            <Option value="Publish">Published</Option>
            <Option value="Draft">Draft</Option>
            <Option value="Trash">Trash</Option>
          </Select>
        </Space>

        <Table
          columns={columns}
          dataSource={articles}
          pagination={{
            total,
            pageSize: limit,
            current: offset / limit + 1,
          }}
          loading={loading}
          onChange={handleTableChange}
        />

        <Modal
          title={currentArticle ? 'Edit Article' : 'Create Article'}
          open={isModalVisible}
          onCancel={handleCancel}
          footer={null}
        >
          <ArticleForm
            article={currentArticle}
            onSuccess={handleOk}
            onCancel={handleCancel}
          />
        </Modal>
      </Space>
    </Card>
  );
};

export default ArticleList;