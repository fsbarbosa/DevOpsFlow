import React, { useState, useEffect } from 'react';
import axios from 'axios';

const CICDComponent = () => {
  const [pipelineData, setPipelineData] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchPipelineData = async () => {
      try {
        const response = await axios.get(process.env.REACT_APP_CI_CD_API_ENDPOINT, {
          headers: { 'Authorization': `Bearer ${process.env.REACT_APP_API_KEY}` },
        });
        setPipelineData(response.data);
        setLoading(false);
      } catch (error) {
        console.error('Error fetching pipeline data:', error);
        setLoading(false);
      }
    };

    fetchPipelineData();
  }, []);

  const triggerDeployment = async () => {
    try {
      await axios.post(process.env.REACT_APP_TRIGGER_DEPLOYMENT_ENDPOINT, {}, {
        headers: { 'Authorization': `Bearer ${process.env.REACT_APP_API_KEY}` },
      });
      setLoading(true);
      setPipelineData([]);
      await fetchPipelineData();
    } catch (error) {
      console.error('Error triggering deployment:', error);
    }
  };

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h2>CI/CD Pipeline Status</h2>
      <button onClick={triggerDeployment}>Trigger New Deployment</button>
      {pipelineData.length > 0 ? (
        pipelineData.map((stage, index) => (
          <div key={index}>
            <h3>{stage.name}</h3>
            <p>Status: {stage.status}</p>
            <p>Progress: {stage.progress}%</p>
          </div>
        ))
      ) : (
        <p>No pipeline data available.</p>
      )}
    </div>
  );
};

export default CICDComponent;