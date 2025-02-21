class LogManager {
  constructor() {
    this.logs = JSON.parse(localStorage.getItem("logs")) || [];
  }

  addLog(action, details) {
    const log = {
      id: Date.now(),
      timestamp: new Date().toISOString(),
      action: action,
      details: details,
    };
    this.logs.unshift(log);
    this.saveToStorage();
  }

  clearLogs() {
    if (!confirm("Are you sure you want to delete all logs?")) return;
    this.logs = [];
    this.saveToStorage();
  }

  saveToStorage() {
    localStorage.setItem("logs", JSON.stringify(this.logs));
  }

  getLogs(filters = {}) {
    let filteredLogs = [...this.logs];

    if (filters.action) {
      filteredLogs = filteredLogs.filter((log) =>
        log.action.includes(filters.action)
      );
    }
    if (filters.startDate) {
      filteredLogs = filteredLogs.filter(
        (log) => new Date(log.timestamp) >= new Date(filters.startDate)
      );
    }
    if (filters.endDate) {
      filteredLogs = filteredLogs.filter(
        (log) => new Date(log.timestamp) <= new Date(filters.endDate)
      );
    }

    return filteredLogs;
  }
}

function displayLogs() {
  const logTableBody = document.querySelector("#logTable tbody");
  logTableBody.innerHTML = "";

  const logs = logManager.getLogs();

  if (logs.length === 0) {
    logTableBody.innerHTML = `<tr><td colspan="2" style="text-align:center;">No logs available</td></tr>`;
    return;
  }

  logs.forEach((log) => {
    const row = document.createElement("tr");
    row.innerHTML = `
            <td>${new Date(log.timestamp).toLocaleString()}</td>
            <td>${log.action} - ${JSON.stringify(log.details)}</td>
        `;
    logTableBody.appendChild(row);
  });
}

const logManager = new LogManager();
