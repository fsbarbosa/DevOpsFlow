const { exec } = require('child_process');
require('dotenv').config();

function triggerPipeline() {
  console.log('Triggering CI/CD pipeline...');
  exec('ci-cd-command-to-trigger-pipeline', (error, stdout, stderr) => {
    if (error) {
      console.error(`Error triggering pipeline: ${error}`);
      sendNotification(`Pipeline Trigger Error: ${error}`);
      return;
    }
    console.log(stdout);
    checkPipelineStatus();
  });
}

function checkPipelineStatus() {
  console.log('Checking pipeline status...');
  setTimeout(() => {
    const pipelineStatus = 'success';
    if (pipelineStatus === 'success') {
      console.log('Pipeline executed successfully.');
      sendNotification('Pipeline executed successfully.');
    } else {
      console.error('Pipeline execution failed.');
      sendNotification('Pipeline execution failed.');
    }
  }, 2000);
}

function sendNotification(message) {
  console.log(`Sending notification: ${message}`);
}

triggerPipeline();