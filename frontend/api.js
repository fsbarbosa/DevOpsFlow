import axios from 'axios';
const API_BASE_URL = process.env.REACT_APP_BACKEND_BASEURL;

let cache = {
  pipelineStatus: { data: null, lastFetch: 0 }
};

const CACHE_DURATION = 5 * 60 * 1000;

const GoBackendService = {
  triggerCICD: async () => {
    try {
      const response = await axios.post(`${API_BASE_URL}/cicd/trigger`);
      console.log('CICD Triggered successfully:', response.data);
      cache.pipelineStatus = { data: null, lastFetch: 0 };
      return response.data;
    } catch (error) {
      console.error('Failed to trigger CICD: ', error);
      throw error;
    }
  },

  getPipelineStatus: async () => {
    const now = new Date().getTime();
    const { data, lastFetch } = cache.pipelineStatus;
    if (data && (now - lastFetch < CACHE_DURATION)) {
      console.log('Returning cached pipeline status');
      return data;
    }
    
    try {
      const response = await axios.get(`${API_BASE_URL}/cicd/status`);
      console.log('Pipeline status received:', response.data);
      cache.pipelineStatus = {
        data: response.data,
        lastFetch: now
      };
      return response.data;
    } catch (error) {
      console.error('Failed to get pipeline status:', error);
      throw error;
    }
  },
  
};

export default GoBackendService;