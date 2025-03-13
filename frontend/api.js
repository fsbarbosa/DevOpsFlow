import axios from 'axios';
const API_BASE_URL = process.env.REACT_APP_BACKEND_BASEURL;
const GoBackendService = {
  triggerCICD: async () => {
    try {
      const response = await axios.post(`${API_BASE_URL}/cicd/trigger`);
      console.log('CICD Triggered successfully:', response.data);
      return response.data;
    } catch (error) {
      console.error('Failed to trigger CICD: ', error);
      throw error;
    }
  },
  getPipelineStatus: async () => {
    try {
      const response = await axios.get(`${API_BASE_URL}/cicd/status`);
      console.log('Pipeline status received:', response.data);
      return response.data;
    } catch (error) {
      console.error('Failed to get pipeline status:', error);
      throw error;
    }
  },
  
};
export default GoBackendService;