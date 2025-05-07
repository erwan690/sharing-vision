import React, { useEffect } from 'react';
import { Form, Input, Select, Button } from 'antd';
import { useNavigate, useParams } from 'react-router-dom';

import { getArticleById, createArticle, updateArticle } from '../services/articleService';

const { TextArea } = Input;
const { Option } = Select;

const ArticleForm = () => {
  const [form] = Form.useForm();
  const navigate = useNavigate();
  const { id } = useParams();
  const isEditMode = Boolean(id);

  useEffect(() => {
    const loadArticle = async () => {
      try {
        const article = await getArticleById(id);
        form.setFieldsValue(article);
      } catch (error) {
        console.error('Failed to load article:', error);
      }
    };

    if (isEditMode) {
      loadArticle();
    }
  }, [id, isEditMode, form]);

  const onFinish = async (values) => {
    try {
      if (isEditMode) {
        await updateArticle(id, values);
      } else {
        await createArticle(values);
      }

      form.resetFields();
      navigate('/all-posts');
    } catch (error) {
      console.error(isEditMode ? 'Error updating article:' : 'Error creating article:', error.message);
    }
  };

  return (
    <Form form={form} layout="vertical" onFinish={onFinish}>
      <Form.Item label="Title" name="title" rules={[{ required: true, min: 20 }]}>
        <Input />
      </Form.Item>

      <Form.Item label="Content" name="content" rules={[{ required: true, min: 200 }]}>
        <TextArea rows={6} />
      </Form.Item>

      <Form.Item label="Category" name="category" rules={[{ required: true, min: 3 }]}>
        <Input />
      </Form.Item>

      <Form.Item label="Status" name="status" rules={[{ required: true }]}>
        <Select>
          <Option value="Publish">Publish</Option>
          <Option value="Draft">Draft</Option>
          <Option value="Trash">Trash</Option>
        </Select>
      </Form.Item>

      <Button type="primary" htmlType="submit">
        {isEditMode ? 'Update' : 'Create'} Article
      </Button>
    </Form>
  );
};

export default ArticleForm;