import api from '../utils/api';


export const createArticle = async (postData) => {
    try {
        const response = await api.post('/article', postData);

        return response.data;
    } catch (error) {
        throw new Error(error.message || 'Failed to create article');
    }
};

export const getArticleById = async (id) => {
    const response = await api.get(`/article/${id}`);
    return response.data;
  };

export const getArticlesByStatus = async (status) => {
    const response = await api.get('/article', {
        params: { status }
    });
    return response.data;
};

export const updateArticle = async (id, data) => {
    const response = await api.put(`/article/${id}`, data);
    return response.data;
};

export const deleteArticle = async (id) => {
    await api.delete(`/article/${id}`);
};