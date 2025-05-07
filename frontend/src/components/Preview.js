import React, { useState, useEffect } from 'react';
import { Table, Pagination } from 'antd';
import { getArticlesByStatus } from '../services/articleService';

const Preview = () => {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(false);
  const [total, setTotal] = useState(0);
  const [limit, setLimit] = useState(10);
  const [offset, setOffset] = useState(0);

  const fetchPublished = async () => {
    setLoading(true);
    try {
      const result = await getArticlesByStatus('Publish');
      setData(result.slice(offset, offset + limit));
      setTotal(result.length);
    } catch (error) {
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchPublished();
  }, [offset, limit]);

  const handlePageChange = (page, pageSize) => {
    setOffset((page - 1) * pageSize);
    setLimit(pageSize);
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
      title: 'Created Date',
      dataIndex: 'created_date',
      key: 'created_date',
    },
  ];

  return (
    <>
      <Table dataSource={data} columns={columns} loading={loading} pagination={false} rowKey="id" />
      <Pagination
        current={offset / limit + 1}
        pageSize={limit}
        total={total}
        onChange={handlePageChange}
        style={{ marginTop: 20, textAlign: 'right' }}
      />
    </>
  );
};

export default Preview;